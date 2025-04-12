Feature: get health
    Scenario: does not allow post method
        When I send "POST" request to "/health"
        Then the response code should be 405
        And the response should match json:
            """
            {
                "message": "Method Not Allowed"
            }
            """
    Scenario: returns Up status
        When I send "GET" request to "/health"
        Then the response code should be 200
        And the response should match json:
            """
            {
                "status": "Up"
            }
            """
