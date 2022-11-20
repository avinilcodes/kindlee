# kindlee

## This app suggest places to live

## Requirements
# postgresql server running on localhost:5432 and having postgres user and kindlee database with password root123

## Usage
# Front-end is not complete yet so we can use it by typing following commands
# go run main.go migrate
# go run main.go start

# One can test this using postman tool
# localhost:8008/get/state
# Request
'''
{
    "personal_residencial_safety":"4",
    "financial_safety":"5",
    "road_safety":"2",
    "workplace_safety":"3",
    "healthcare_facilities":"5",
    "monthly_rent_spare":"1500"
}
'''
Response
'''
{
    "best_suits": "Minnesota,Utah,Vermont"
}
'''
