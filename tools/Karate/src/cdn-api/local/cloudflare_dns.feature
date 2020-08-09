Feature: Test Cloudflare API DNS
  Background:
    * url 'http://127.0.0.1:8080'

  Scenario: List DNS Record
    * def data = read('karate_config_data.json')

    Given url data["scenario"]["cloudflare"]["list_dns_record"]["url"]
    And headers data["scenario"]["cloudflare"]["list_dns_record"]["header"]
    And request data["scenario"]["cloudflare"]["list_dns_record"]["body"]
    When method post
    Then status 200

  Scenario: Create DNS Record
    * def data = read('karate_config_data.json')

    Given url data["scenario"]["cloudflare"]["create_dns_record"]["url"]
    And headers data["scenario"]["cloudflare"]["create_dns_record"]["header"]
    And request data["scenario"]["cloudflare"]["create_dns_record"]["body"]
    When method post
    Then status 200

  Scenario: DNS Record Details
    * def data = read('karate_config_data.json')

    Given url data["scenario"]["cloudflare"]["dns_record_details"]["url"]
    And headers data["scenario"]["cloudflare"]["dns_record_details"]["header"]
    And request data["scenario"]["cloudflare"]["dns_record_details"]["body"]
    When method post
    Then status 200

  Scenario: Delete DNS Record
    * def data = read('karate_config_data.json')

    Given url data["scenario"]["cloudflare"]["delete_dns_record"]["url"]
    And headers data["scenario"]["cloudflare"]["delete_dns_record"]["header"]
    And request data["scenario"]["cloudflare"]["delete_dns_record"]["body"]
    When method post
    Then status 200
