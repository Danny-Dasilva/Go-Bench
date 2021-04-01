for i in {1..1000};do curl -s -w "%{time_total}\n" -o /dev/null localhost:8080; done
