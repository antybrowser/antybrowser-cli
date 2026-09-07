{ lib
, buildGoModule
, fetchFromGitHub
, installShellFiles
}:

buildGoModule rec {
  pname = "antybrowser-cli";
  version = "1.0.0";

  src = fetchFromGitHub {
    owner = "antybrowser";
    repo = "antybrowser-cli";
    rev = "cli-v${version}";
    hash = "PLACEHOLDER_SHA256";
  };

  vendorHash = "PLACEHOLDER_VENDOR_HASH";

  subPackages = [ "cmd/antybrowser" ];

  ldflags = [
    "-s"
    "-w"
    "-X main.version=${version}"
  ];

  postInstall = ''
    installShellCompletion --cmd antybrowser \
      --bash <($out/bin/antybrowser completion bash) \
      --fish <($out/bin/antybrowser completion fish) \
      --zsh <($out/bin/antybrowser completion zsh)
  '';

  meta = with lib; {
    description = "CLI tool for managing Antybrowser profiles, proxies, and automations";
    homepage = "https://antybrowser.com";
    license = licenses.mit;
    maintainers = [ ];
    mainProgram = "antybrowser";
    platforms = platforms.unix ++ platforms.windows;
  };
}
