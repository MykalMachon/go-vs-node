import { Hono } from 'hono'
import { validator } from 'hono/validator'

import { serve } from '@hono/node-server'

import sqlite3 from 'sqlite3'
import { open } from 'sqlite'


const app = new Hono()

// create a new SQLite database if it doesn't exist
const db = await open({
  filename: 'posts.db',
  driver: sqlite3.Database
})

// create a posts table if it doesn't exist
await db.run(`
  CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    content TEXT
  )
`)

// basic route
app.get('/', (c) => {
  return c.body('Hello World')
});


// a route to create an anonymous post 
app.post('/posts', validator('form', (value, c) => {
  const title = value['title']
  const content = value['content']
  if (!title || !content) {
    return c.text('invalid: title and content are required', 400)
  }
  return {
    title,
    content
  }
}), async (c) => {
  const { title, content } = c.req.valid('form')
  const result = await db.run('INSERT INTO posts (title, content) VALUES (?, ?)', title, content)
  if (result.changes === 1) {
    return c.text('Post created', 201)
  } else {
    return c.text('An error occurred', 500)
  }
})

// a route to get the recent 10 posts 
app.get('/posts', async (c) => {
  const results = await db.all('SELECT * FROM posts ORDER BY id DESC LIMIT 10')
  return c.json(results)
})

console.log('Starting server on http://localhost:3000')
serve({
  fetch: app.fetch,
  port: 3000,
})