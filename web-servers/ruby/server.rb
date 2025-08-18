require 'webrick'

PORT = 8080

server = WEBrick::HTTPServer.new(
  Port: PORT
)

server.mount_proc '/' do |req, res|
  if req.path == '/health'
    res.status = 200
    res['Content-Type'] = 'text/plain'
    res.body = 'ok'
  end
end

trap('INT') { server.shutdown }

puts "Ruby server running on http://localhost:#{PORT}"
server.start
