#!/usr/bin/env python3
"""Installs latest version to system"""
import platform
from urllib.request import urlopen, urlretrieve, Request
from os import system
from json import loads

API_URL = "https://api.github.com/repos/mishamyrt/vpn-dns/releases/latest"
LOCAL_PATH = "./vpn-dns"
SYSTEM_PATH = "/usr/local/bin/vpn-dns"

def get_latest_release():
    """Fetches latest release information from GitHub"""
    req = Request(API_URL, headers={
        "Accept": "application/vnd.github+json",
        "X-GitHub-Api-Version": "2022-11-28"
    })
    with urlopen(req) as response:
        return loads(response.read().decode())

def find_binary(assets, arch):
    """Installs binary asset for arch"""
    for asset in assets:
        if asset["name"].endswith(arch):
            return asset
    return None

def install():
    """Installs vpn-dns to system"""
    if platform.system() != 'Darwin':
        print("Only macOS is supported at the moment")
        exit(1)
    print("Getting release...")
    release = get_latest_release()
    arch = platform.machine().lower()
    binary_asset = find_binary(release["assets"], arch)
    if binary_asset is None:
        print(f"Binary for your arch '{arch}' is not found :(")
        exit(1)
    print(f"Downloading {release['tag_name']} for {arch}")
    print(binary_asset["browser_download_url"])
    urlretrieve(binary_asset["browser_download_url"], "./vpn-dns")
    print("Installing to system")
    system(f"chmod +x {LOCAL_PATH}")
    system(f"sudo cp -f {LOCAL_PATH} {SYSTEM_PATH}")
    print("Done")

install()
