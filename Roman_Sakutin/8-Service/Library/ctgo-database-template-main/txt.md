для теста

cd Roman_Sakutin/8-Service/Library/ctgo-database-template-main

docker compose up -d postgres

make build

export POSTGRES_HOST=127.0.0.1
export POSTGRES_PORT=5432
export POSTGRES_DB=library
export POSTGRES_USER=library
export POSTGRES_PASSWORD=1234567
export POSTGRES_MAX_CONN=10

для локи

docker compose up library promtail loki grafana

для пост аутбокс сервера

python3 -c "
from http.server import BaseHTTPRequestHandler, HTTPServer

class H(BaseHTTPRequestHandler):
    def do_POST(self):
        length = int(self.headers.get('Content-Length', 0))
        print('received:', self.rfile.read(length))
        self.send_response(200)
        self.end_headers()

    def log_message(self, *args):
        pass

HTTPServer(('0.0.0.0', 8099), H).serve_forever()
"
