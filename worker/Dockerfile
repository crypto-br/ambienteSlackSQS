FROM public.ecr.aws/lambda/python:3.9

RUN python -m pip install --upgrade pip

# Instalando dependencias do python
RUN python -m pip install boto3
RUN python -m pip install localstack_client
RUN python -m pip install requests
RUN python -m pip install flask

WORKDIR /app

ADD handler.py ./

ENTRYPOINT [ "python" ]

EXPOSE 8082

CMD ["/app/handler.py"]
