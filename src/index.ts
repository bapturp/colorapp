import express from 'express'

const app = express()

app.use(express.json())

app.get('/', (_, res) => {
  const color = process.env.COLOR || 'unknown'
  res.json({ color })
})

let ready = false
app.get('/healthz/readiness', (_, res) => {
  if (!ready) {
    res.status(503).json({ status: 'NOT READY' })
  }
  res.json({ status: 'OK' })
})

app.get('/healthz/liveness', (_, res) => {
  res.json({ status: 'OK' })
})

const port = process.env.PORT || 8080
const server = app.listen(port, () => {
  ready = true
  console.log(`Server is running on port ${port}`)
})

process.on('SIGTERM', () => {
  console.log('SIGTERM received, closing server')
  server.close(() => {
    console.log('server terminated')
  })
})
