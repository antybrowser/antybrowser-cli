class AntybrowserCli < Formula
  desc "CLI for Antybrowser anti-detect browser - manage profiles, proxies, extensions, and automations"
  homepage "https://antybrowser.com"
  version "1.0.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-darwin-arm64.tar.gz"
      sha256 "4a55d93388be5227f0f596f745057ef8418a7de70242d007017f1c7cd56dc889"
    else
      url "https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-darwin-amd64.tar.gz"
      sha256 "daa594ff4514992e456a6bdf1a04b37c21bbce835aa854bf1f8f19f34fa7e861"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-linux-arm64.tar.gz"
      sha256 "d1d16f35473b08261f4509d942dbaa451bbbb9cd6a2cac2f5fab92b46ce98445"
    else
      url "https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-linux-amd64.tar.gz"
      sha256 "afaa12f4ad4c7fe9d2c8302818db11e26b3a0e66a82937e3785403a05e808a8b"
    end
  end

  def install
    bin.install "antybrowser"
  end

  test do
    system "#{bin}/antybrowser", "--version"
  end
end
