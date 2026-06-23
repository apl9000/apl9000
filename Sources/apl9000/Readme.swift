import Foundation
import Stencil

/// The full data model handed to the renderer, mirroring the Go `ProfileData`.
struct ProfileData {
    let timestamp: String
    let website: Website
    let rates: Rates
}

enum RenderError: Error {
    case templateNotFound
}

/// Renders the profile README HTML from `Templates/profile.stencil`.
///
/// The template is bundled as a package resource and rendered with Stencil,
/// preserving the template/code separation of the original Go `text/template`.
/// Stencil does not HTML-escape by default, matching `text/template` semantics
/// (the website description contains raw anchor markup).
func renderReadme(_ data: ProfileData) throws -> String {
    guard let url = Bundle.module.url(
        forResource: "profile",
        withExtension: "stencil",
        subdirectory: "Templates"
    ) else {
        throw RenderError.templateNotFound
    }

    let templateString = try String(contentsOf: url, encoding: .utf8)

    let context: [String: Any] = [
        "timestamp": data.timestamp,
        "website": [
            "url": data.website.url,
            "description": data.website.description,
            "blogPosts": data.website.blogPosts.map { post in
                [
                    "pubDate": post.publishedAt,
                    "slug": post.slug,
                    "title": post.title,
                    "summary": post.summary,
                ]
            },
            "socials": data.website.socials.map { social in
                ["href": social.href, "name": social.name]
            },
            "quoteOfTheDay": [
                "text": data.website.quoteOfTheDay.text,
                "author": data.website.quoteOfTheDay.author,
            ],
        ],
        "rates": [
            "cad": data.rates.cad,
            "usd": data.rates.usd,
            "eur": data.rates.eur,
            "gbp": data.rates.gbp,
            "jpy": data.rates.jpy,
            "mxn": data.rates.mxn,
            "aud": data.rates.aud,
        ],
    ]

    let environment = Environment()
    return try environment.renderTemplate(string: templateString, context: context)
}
