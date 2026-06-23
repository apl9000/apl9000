// swift-tools-version:5.9
import PackageDescription

let package = Package(
    name: "apl9000",
    platforms: [
        .macOS(.v13)
    ],
    targets: [
        .executableTarget(
            name: "apl9000",
            path: "Sources/apl9000"
        )
    ]
)
