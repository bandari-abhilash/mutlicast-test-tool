const server = Bun.serve({
    port:Bun.env.port,
    fetch(req){
        return new Response("Start!")
    }
})

console.log(`Server up and running on port ${server.port}`)