const express = require('express');
const client = require('prom-client');
const { pool } = require('./db');

const app = express();
app.use(express.json());

const register = new client.Registry();
client.collectDefaultMetrics({ register });

const httpRequestDuration = new client.Histogram({
  name: 'http_request_duration_seconds',
  help: 'HTTP request duration in seconds',
  labelNames: ['method', 'route', 'status_code'],
  registers: [register],
});

app.use((req, res, next) => {
  const end = httpRequestDuration.startTimer();
  res.on('finish', () => {
    end({ method: req.method, route: req.path, status_code: res.statusCode });
  });
  next();
});

app.get('/healthz', (req, res) => {
  res.json({ status: 'ok' });
});

app.get('/items', async (req, res) => {
  const result = await pool.query('SELECT id, value, created_at FROM items ORDER BY id DESC LIMIT 100');
  res.json(result.rows);
});

app.post('/items', async (req, res) => {
  const { value } = req.body;
  const result = await pool.query('INSERT INTO items (value) VALUES ($1) RETURNING id, value, created_at', [value]);
  res.status(201).json(result.rows[0]);
});

app.put('/items/:id', async (req, res) => {
  const { value } = req.body;
  const result = await pool.query(
    'UPDATE items SET value = $1 WHERE id = $2 RETURNING id, value, created_at',
    [value, req.params.id]
  );
  if (result.rowCount === 0) {
    return res.status(404).json({ error: 'not found' });
  }
  res.json(result.rows[0]);
});

app.delete('/items/:id', async (req, res) => {
  const result = await pool.query('DELETE FROM items WHERE id = $1', [req.params.id]);
  if (result.rowCount === 0) {
    return res.status(404).json({ error: 'not found' });
  }
  res.status(204).end();
});

app.post('/reset', async (req, res) => {
  await pool.query('TRUNCATE TABLE items RESTART IDENTITY');
  res.json({ status: 'reset' });
});

app.get('/metrics', async (req, res) => {
  res.set('Content-Type', register.contentType);
  res.end(await register.metrics());
});

const port = Number(process.env.PORT || 4000);
app.listen(port, () => {
  console.log(`node-express benchmark app listening on :${port}`);
});
