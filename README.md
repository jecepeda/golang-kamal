# Golang Kamal

This repository servers as an example of how to run Golang applications using
[Kamal](https://kamal-deploy.org/)

## Installation

For this app to be deployed, you need to run in your terminal

```bash
gem install kamal
```

Afterwards, go to [the deploy file](config/deploy.yml), and change it according to your needs. You would probably have to change `<username>` for the one you have in your Docker registry.

Then, create a `.env` file and paste your registry token in a variable called `KAMAL_REGISTRY_PASSWORD`.

Finally, you can deploy your application using

`kamal setup`
