getFlagVal() {
    local name=$1
    local default=$2
    local count=0

    for arg in "$@"; do
        if [[ $arg == "--$name="* ]]; then
            local value=${arg#*=}
            echo "$value"
            count=1
            break
        fi
    done

    if [[ $count -eq 0 ]]; then
        echo "$default"
    fi
}

case $1 in
dev)
    eval $(awk '!/^#/ && NF > 0 {print "export "  $1}' .env.dev)
    m=$(getFlagVal "m" false $@)

    if [[ $m == "true" ]]; then
        make run_migrate
    else
        make run
    fi
    ;;

docker)
    eval $(awk '!/^#/ && NF > 0 {print "export "  $1}' .env.docker) && docker compose up -d --build
    ;;

prod)
    eval $(awk '!/^#/ && NF > 0 {print "export "  $1}' .env.prod) && make deploy
    ;;

set)
    eval $(awk '!/^#/ && NF > 0 {system("flyctl secrets set " $1)}' .env.prod)
    ;;
*)
    echo -n "Unknown command"
    ;;
esac
