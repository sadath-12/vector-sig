send_events() {
    for i in $(seq 1 20)
    do
        curl -s http://localhost:8081/elastic/_bulk --data-binary "@2kevents.json" -o res.txt
        if [ $? -ne 0 ]; then
            post_event "install_failed" "send_events: Failed to send sample log dataset to http://localhost:8081/elastic/_bulk"
            print_error_and_exit "Failed to send sample log dataset"
        fi
    done
    echo "\n Sample log dataset sent successfully"
}

send_events