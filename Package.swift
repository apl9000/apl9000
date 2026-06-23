// swift-tools-version:5.9
import PackageDescription

let package = Package(
    name: "apl9000",
    platforms: [
        .macOS(.v13)
    ],
    dependencies: [
        .package(url: "https://github.com/stencilproject/Stencil.git", from: "0.15.1")
    ],
    targets: [
        .executableTarget(
            name: "apl9000",
            dependencies: ["Stencil"],
            path: "Sources/apl9000",
            resources: [
                .copy("Templates")
            ]
        )
    ]
)
