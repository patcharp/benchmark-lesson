## Example project to run test

ประกอบไปด้วย Framework ที่ใช้กันบ่อยๆ อยู่ 5 profiles
1. ElysiaJS กับ Bun
2. ExpressJS กับ NPM
3. Go Echo
4. Go Fiber
5. Go Gin
6. Python Flask

วิธีการ Run ก็คือรันโดยใช้คำสั่ง `docker compose --profile <PROFILE> up --build` โดย profile ก็จะอยู่ที่เราจะเลือก Framework ตัวไหน อย่างเช่น
1. ElysiaJS -> elysia
2. ExpressJS -> express
3. Go Echo -> go-echo
4. Go Fiber -> go-fiber
5. Go Gin -> go-gin
6. Python Flask -> flask

ตัวอย่างเช่น ต้องการ Start ElysiaJS ก็จะรันคำสั่งดังนี้

`docker compose --profile elysia up --build`
