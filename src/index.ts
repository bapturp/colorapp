import express from 'express';

const app = express();

app.use(express.json());

app.get('/', (_, res) => {
  const color = process.env.COLOR || 'unknown';
  res.json({ color });
});

let ready = false;
app.get('/health/ready', (_, res) => {
  ready
    ? res.json({ status: 'OK' })
    : res.status(503).json({ status: 'NOT READY' });
});

app.get('/health/live', (_, res) => {
  res.json({ status: 'OK' });
});

const port = process.env.PORT || 8080;
const server = app.listen(port, () => {
  ready = true;
  console.log(`Server is running on port ${port}`);
});

process.on('SIGTERM', () => {
  server.close(() => {
    console.log('Process terminated');
  });
});
