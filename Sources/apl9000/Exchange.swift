import Foundation

private struct ExchangeRateResponse: Decodable {
    let rates: [String: Double]
}

struct Rates {
    var cad = ""
    var usd = ""
    var eur = ""
    var gbp = ""
    var jpy = ""
    var mxn = ""
    var aud = ""
}

func getRates() async -> Rates {
    // https://www.exchangerate-api.com/docs/free
    let uri = "https://open.er-api.com/v6/latest/CAD"

    guard let response = await fetchJSON(uri, as: ExchangeRateResponse.self) else {
        return Rates()
    }

    func format(_ code: String) -> String {
        String(format: "%.2f", response.rates[code] ?? 0)
    }

    return Rates(
        cad: format("CAD"),
        usd: format("USD"),
        eur: format("EUR"),
        gbp: format("GBP"),
        jpy: format("JPY"),
        mxn: format("MXN"),
        aud: format("AUD")
    )
}
