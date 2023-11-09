FROM python:3.9-slim

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt


COPY replica_monitor.py .

# Prometheus client is using
EXPOSE 8090


CMD [ "python", "./replica_monitor.py" ]
