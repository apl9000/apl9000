import Foundation

/// The full data model handed to the renderer, mirroring the Go `ProfileData`.
struct ProfileData {
    let timestamp: String
    let website: Website
    let weatherForecast: WeatherForecast
    let rates: Rates
}

/// Renders the profile README HTML, faithfully reproducing the output of the
/// original Go `text/template` in `template/index.html`.
///
/// Note: like Go's `text/template` (as opposed to `html/template`), values are
/// interpolated verbatim without HTML escaping.
func renderReadme(_ data: ProfileData) -> String {
    var html = ""

    // Description paragraphs
    for line in data.website.description {
        html += "\n<p>\(line)</p>\n"
    }

    // Blog
    html += "<h3>Blog</h3>\n"
    for post in data.website.blogPosts {
        html += " \(post.publishedAt)\n"
        html += "<br />\n"
        html += "<a href=\"\(data.website.url)blog/\(post.slug)\" target=\"_blank\">\n"
        html += "\(post.title) - \(post.summary)</a>\n"
        html += "<br />\n"
        html += "<br />\n"
    }

    // Quote of the day
    html += "<h3>Quote of the day :)</h3>\n"
    html += "<text\n  >\(data.website.quoteOfTheDay.text)<br />\n"
    html += "  —\(data.website.quoteOfTheDay.author)</text\n>\n"

    // Socials
    html += "<h3>Where to find me?</h3>\n"
    html += "<p>\n"
    for social in data.website.socials {
        html += "  <a href=\"\(social.href)\" target=\"_blank\">\(social.name)</a><br />\n"
    }
    html += "</p>\n"

    // Exchange rates table
    html += exchangeRatesTable(data.rates)

    // Footer
    html += "<br />\n"
    html += "<img\n"
    html += "  alt=\"README Update\"\n"
    html += "  src=\"https://github.com/apl9000/apl9000/actions/workflows/readme_update.yaml/badge.svg\"\n"
    html += "/>\n"
    html += "<p>Last updated: \(data.timestamp)</p>\n"

    return html
}

private func exchangeRatesTable(_ rates: Rates) -> String {
    let rows: [(flag: String, width: Int, code: String, rate: String)] = [
        ("https://upload.wikimedia.org/wikipedia/commons/thumb/d/d9/Flag_of_Canada_%28Pantone%29.svg/2880px-Flag_of_Canada_%28Pantone%29.svg.png", 68, "CAD", rates.cad),
        ("https://upload.wikimedia.org/wikipedia/en/thumb/a/a4/Flag_of_the_United_States.svg/1600px-Flag_of_the_United_States.svg.png?20151118161041", 66, "USD", rates.usd),
        ("https://upload.wikimedia.org/wikipedia/commons/thumb/f/fc/Flag_of_Mexico.svg/2560px-Flag_of_Mexico.svg.png", 66, "MXN", rates.mxn),
        ("https://upload.wikimedia.org/wikipedia/commons/thumb/b/b7/Flag_of_Europe.svg/2560px-Flag_of_Europe.svg.png", 66, "EUR", rates.eur),
        ("https://upload.wikimedia.org/wikipedia/en/thumb/a/ae/Flag_of_the_United_Kingdom.svg/1920px-Flag_of_the_United_Kingdom.svg.png", 66, "GBP", rates.gbp),
        ("https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Flag_of_Australia_%28converted%29.svg/2560px-Flag_of_Australia_%28converted%29.svg.png", 66, "AUD", rates.aud),
        ("https://upload.wikimedia.org/wikipedia/en/thumb/9/9e/Flag_of_Japan.svg/1920px-Flag_of_Japan.svg.png", 66, "JPY", rates.jpy),
    ]

    var table = "<h3>Exchange Rates</h3>\n<table>\n"
    table += "  <tr>\n    <th></th>\n    <th>Currency</th>\n    <th>Rate</th>\n  </tr>\n"
    for row in rows {
        table += "  <tr>\n"
        table += "    <td>\n"
        table += "      <img\n        alt=\"\(row.code) Flag\"\n        width=\"\(row.width)\"\n        src=\"\(row.flag)\"\n      />\n"
        table += "    </td>\n"
        table += "    <td>\(row.code)</td>\n"
        table += "    <td>\(row.rate)</td>\n"
        table += "  </tr>\n"
    }
    table += "</table>\n"
    return table
}
