class AntybrowserCli < Formula
  desc "CLI tool for managing Antybrowser profiles, proxies, and automations"
  homepage "https://antybrowser.com"
  version "1.0.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-darwin-arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256"
    else
      url "https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-darwin-amd64.tar.gz"
      sha256 "PLACEHOLDER_SHA256"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-linux-arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256"
    else
      url "https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-linux-amd64.tar.gz"
      sha256 "PLACEHOLDER_SHA256"
    end
  end

  def install
    bin.install "antybrowser"
  end

  test do
    system "#{bin}/antybrowser", "--version"
  end
end
