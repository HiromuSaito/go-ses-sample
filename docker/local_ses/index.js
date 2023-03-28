import server from 'aws-ses-v2-local'
let s = await server.default({ port: 8005 })
console.log('aws-ses-v2-local: server up and running')