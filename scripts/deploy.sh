ssh -t "$TRUCO_SERVER" <<-'ENDSSH'
    cd truco
    git pull
    go build -o truco-executable
    cd frontend
    npm install
    npm run build
    sudo service truco restart
ENDSSH

echo "Website deployed 🎉"
xdg-open "https://bkach.com"
