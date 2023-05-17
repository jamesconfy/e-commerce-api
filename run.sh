# Running with docker 
export $(cat .env.docker | xargs) && docker compose up -d --build
# Running for development
export $(cat .env.dev | xargs) && make run
# Running in production, using flyctl
export $(cat .env.prod | xargs) && make deploy
# Deploy secrets to fly
awk '{system("flyctl secrets set " $1)}'