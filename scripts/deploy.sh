ssh -t "$TRUCO_SERVER" <<-'ENDSSH'
    cd truco
    git pull
    go build -o truco-executable
    sudo service truco restart
ENDSSH

echo "Website deployed 🎉"