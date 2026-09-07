cask "antybrowser" do
  version "2.0.9"
  sha256 "PLACEHOLDER_SHA256"

  url "https://github.com/antybrowser/antybrowser/releases/download/v#{version}/Antybrowser-#{version}-mac-universal.dmg"
  name "Antybrowser"
  desc "Anti-detect browser for multi-account management"
  homepage "https://antybrowser.com"

  livecheck do
    url :url
    strategy :github_latest
  end

  app "Antybrowser.app"

  zap trash: [
    "~/Library/Application Support/Antybrowser",
    "~/Library/Preferences/com.antybrowser.app.plist",
    "~/Library/Saved Application State/com.antybrowser.app.savedState",
    "~/Library/Caches/Antybrowser",
  ]
end
