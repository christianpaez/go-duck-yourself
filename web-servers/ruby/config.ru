require 'rack'

files = Rack::Files.new('public')

app = lambda do |env|
  if env['PATH_INFO'] == '/health'
    [200, { 'content-type' => 'text/plain' }, ['ok']]
  else
    files.call(env)
  end
end

run app
