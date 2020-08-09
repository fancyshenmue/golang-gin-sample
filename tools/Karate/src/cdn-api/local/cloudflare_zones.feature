Feature: Test Cloudflare API Zone
  Background:
    * url 'http://127.0.0.1:8080'

  Scenario: List Zones
    * def data = read('karate_config_data.json')

    Given url data["scenario"]["cloudflare"]["list_zone"]["url"]
    And headers data["scenario"]["cloudflare"]["list_zone"]["header"]
    And request data["scenario"]["cloudflare"]["list_zone"]["body"]
    When method post
    Then status 200

  Scenario: Zones Details
    * def data = read('karate_config_data.json')

    Given url data["scenario"]["cloudflare"]["zone_details"]["url"]
    And headers data["scenario"]["cloudflare"]["zone_details"]["header"]
    And request data["scenario"]["cloudflare"]["zone_details"]["body"]
    When method post
    Then status 200
