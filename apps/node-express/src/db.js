const { Pool } = require('pg');

const pool = new Pool({
  host: process.env.PGHOST || 'localhost',
  port: Number(process.env.PGPORT || 5432),
  user: process.env.PGUSER || 'bench',
  password: process.env.PGPASSWORD,
  database: process.env.PGDATABASE || 'bench',
});

module.exports = { pool };
