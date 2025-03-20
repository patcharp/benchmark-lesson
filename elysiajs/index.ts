import { Elysia } from 'elysia'

const app = new Elysia()
	.get('/', () => 'elysiajs')
	.listen(3000)

console.log(`🦊 Elysia is running at ${app.server?.port}`)
