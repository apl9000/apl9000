import Foundation

// Subset of the api.weather.gov forecast response that we actually use.
private struct WeatherResponse: Decodable {
    let properties: WeatherProperties
}

private struct WeatherProperties: Decodable {
    let periods: [WeatherPeriod]
}

private struct WeatherPeriod: Decodable {
    let temperature: Int
    let detailedForecast: String
}

struct Temperature {
    let celsius: Int
    let fahrenheit: Int
}

struct WeatherForecast {
    let summary: String
    let temperature: Temperature

    static let empty = WeatherForecast(
        summary: "",
        temperature: Temperature(celsius: 0, fahrenheit: 0)
    )
}

func fahrenheitToCelsius(_ f: Int) -> Int {
    (f - 32) * 5 / 9
}

func getWeatherForecast() async -> WeatherForecast {
    // https://www.weather.gov/documentation/services-web-api
    let nycURI = "https://api.weather.gov/gridpoints/OKX/33,35/forecast"

    guard let response = await fetchJSON(nycURI, as: WeatherResponse.self),
          let first = response.properties.periods.first else {
        return .empty
    }

    return WeatherForecast(
        summary: first.detailedForecast,
        temperature: Temperature(
            celsius: fahrenheitToCelsius(first.temperature),
            fahrenheit: first.temperature
        )
    )
}
