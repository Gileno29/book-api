
# 🌐 API books 
<div align="center">
  <img src="./frontend/public/banner.png" alt="Logo do Projeto" width="200"/>
</div>

<p align="center">
  <a href="https://github.com/Gileno29/agenda">
    <img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/guedes-jr/django_next_auth">
  </a>
  <a href="https://github.com/Gileno29/agenda/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/guedes-jr/django_next_auth">
  </a>
  <a href="https://github.com/Gileno29/agenda/network">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/guedes-jr/django_next_auth">
  </a>
  <a href="https://github.com/Gileno29/agenda/stargazers">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/guedes-jr/django_next_auth">
  </a>
  <a href="https://github.com/Gileno29/agenda/blob/main/LICENSE">
    <img alt="GitHub license" src="https://img.shields.io/github/license/guedes-jr/django_next_auth">
  </a>
</p>

## 📝 Sumário

- [Sobre o Projeto](#sobre-o-projeto)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Funcionalidades](#funcionalidades)
- [Requisitos](#requisitos)
- [Instalação](#instalação)
- [Uso](#uso)
- [Scripts Disponíveis](#scripts-disponíveis)
- [Estrutura de Pastas](#estrutura-de-pastas)
- [Contribuindo](#contribuindo)
- [Licença](#licença)
- [Contato](#contato)

## 🛠️ Sobre o Projeto

Este é um projeto full-stack que combina Django para o back-end e Next.js para o front-end. A aplicação visa fornecer uma plataforma robusta para novos projetos.

## 🧰 Tecnologias Utilizadas

- [Django](https://www.djangoproject.com/) - Back-end framework
- [Docker](https://www.docker.com/) - Deploy
- [SQLitle](https://www.sqlite.org/) - Banco de dados
- [Nginx](https://nginx.org/en/) - Servidor Web

## ✨ Funcionalidades

- Inserção de evendos
- Atualização
- Remoção
- Listagem

## 📋 Requisitos

- Python 3.10
- Docker
- Django

## 🚀 Instalação

### Clonando o Repositório

```bash
git clone https://github.com/Gileno29/agenda

cd agenda
```

### Configurando o Back-end (Django)

```bash
# Criar ambiente virtual
python3 -m venv agenda-venv

# Ativar ambiente virtual
source agenda-venv/bin/activate  # No Windows use `venv\Scripts\activate`

# Instalar dependências necessárias para execução do projeto
pip install -r requirements.txt

#Rodar a aplicação em modo de desenvolvimento

python manage.py runserver
```


### Executando a Aplicação em produção

```bash
sudo docker-compose up 
  #execute com a flag -d para executar em segundo plano


#Parar aplicação

docker-compose down -v
```
## 📁 Estrutura de Pastas

```plaintext
├── agenda
│   ├── agenda
│   │   ├── asgi.py
│   │   ├── __init__.py
│   │   ├── settings.py
│   │   ├── urls.py
│   │   └── wsgi.py
│   ├── core
│   │   ├── admin.py
│   │   ├── apps.py
│   │   ├── __init__.py
│   │   ├── models.py
│   │   ├── tests.py
│   │   └── views.py
│   ├── db.sqlite3
│   ├── manage.py
│   ├── static
│   │   └── css
│   │       ├── agenda.css
│   │       ├── evento.css
│   │       └── login_style.css
│   └── templates
│       ├── agenda.html
│       ├── evento.html
│       ├── login.html
│       ├── model-footer.html
│       ├── model-header.html
│       └── model-page.html
├── db.sqlite3
├── docker-compose.yml
├── Dockerfile
├── entrypoint.sh
├── nginx
│   ├── Dockerfile
│   └── nginx.conf
├── readme.md
└── requirements.txt
```
> Comando utilizado para mostrar a estrutura de dados `tree  -I '__pycache__' -I 'migrations' -I 'agenda_env'`.

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

1. Faça um fork do projeto
2. Crie uma nova branch (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas alterações (`git commit -m 'Adiciona nova funcionalidade'`)
4. Faça o push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 📧 Contato

👤 **Seu Nome**

- Github: [@Gileno29](https://github.com/Gileno29/agenda)
- LinkedIn: [Gileno Duarte](https://www.linkedin.com/in/gileno-cordeiro-duarte-75913a164/)
- Email: gilenoduarte.jobs@gmail.com
---

Desenvolvido com profissionalismo por [Gileno Duarte](https://github.com/Gileno29/agenda) 🤖.