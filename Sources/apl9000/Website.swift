import Foundation

// Shape of the https://www.apl.directory/api response.
private struct WebsiteResponseData: Decodable {
    let metadata: Metadata
    let quoteOfTheDay: QuoteOfTheDay
}

private struct Metadata: Decodable {
    let url: String
    let title: String
    let description: [String]
    let socials: [Social]
}

struct Social: Decodable {
    let name: String
    let href: String
}

struct QuoteOfTheDay: Decodable {
    let text: String
    let author: String

    static let empty = QuoteOfTheDay(text: "", author: "")
}

struct Website {
    let title: String
    let description: [String]
    let url: String
    let socials: [Social]
    let quoteOfTheDay: QuoteOfTheDay

    static let empty = Website(
        title: "",
        description: [],
        url: "",
        socials: [],
        quoteOfTheDay: .empty
    )
}

func getWebsiteData() async -> Website {
    let url = "https://www.apl.directory/api"

    guard let data = await fetchJSON(url, as: WebsiteResponseData.self) else {
        return .empty
    }

    return Website(
        title: data.metadata.title,
        description: data.metadata.description,
        url: data.metadata.url,
        socials: data.metadata.socials,
        quoteOfTheDay: data.quoteOfTheDay
    )
}
