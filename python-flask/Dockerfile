FROM python:3
WORKDIR /app

COPY . .
RUN pip install -r requirements.txt

CMD [ "gunicorn", "-w", "4", "-b", "0.0.0.0:3000", "app:app" ]
