ssh -t "$TRUCO_SERVER" <<-'ENDSSH'
    cd truco
    git pull
    go build -o truco-executable
    cd frontend
    yarn install
    yarn build
    sudo service truco restart
ENDSSH

echo "Website deployed 🎉"
xdg-open "https://bkach.com"
