import requests
import json
import sys

def geocode_address(address):
    """Retourne les coordonnées (lat, lon) d'une adresse via Nominatim."""
    url = "https://nominatim.openstreetmap.org/search"
    params = {
        "q": address,
        "format": "jsonv2",
        "limit": 1,
        "addressdetails": 1
    }
    headers = {"User-Agent": "KestrelApp/1.0 (contact@example.com)"}
    r = requests.get(url, params=params, headers=headers)
    r.raise_for_status()
    data = r.json()
    if not data:
        raise ValueError("Adresse introuvable.")
    return float(data[0]["lat"]), float(data[0]["lon"])

def search_nearby(lat, lon, radius=300):
    """Interroge Overpass API pour obtenir les points d'intérêt autour d'une position."""
    overpass_url = "https://overpass-api.de/api/interpreter"
    query = f"""
    [out:json][timeout:25];
    (
      node["amenity"](around:{radius},{lat},{lon});
      node["shop"](around:{radius},{lat},{lon});
    );
    out body;
    """
    r = requests.get(overpass_url, params={'data': query})
    r.raise_for_status()
    return r.json()["elements"]

def extract_info(elements, address):
    """Filtre les lieux pertinents et extrait les informations intéressantes."""
    infos = []
    for el in elements:
        tags = el.get("tags", {})
        if not tags:
            continue

        # On garde si le nom correspond à l'adresse ou si c'est un commerce identifiable
        if "name" in tags:
            info = {
                "nom": tags.get("name"),
                "type": tags.get("amenity") or tags.get("shop"),
                "téléphone": tags.get("phone") or tags.get("contact:phone"),
                "email": tags.get("email") or tags.get("contact:email"),
                "site": tags.get("website") or tags.get("contact:website"),
                "horaires": tags.get("opening_hours"),
                "adresse_OSM": f"{tags.get('addr:housenumber', '')} {tags.get('addr:street', '')}".strip(),
                "autres_tags": {k: v for k, v in tags.items() if k not in [
                    "name", "phone", "contact:phone", "email", "contact:email",
                    "website", "contact:website", "opening_hours",
                    "addr:housenumber", "addr:street"
                ]}
            }
            infos.append(info)
    return infos

def main():
    if len(sys.argv) < 2:
        print("Usage : python find_place.py '<adresse>'")
        sys.exit(1)

    address = sys.argv[1]
    print(f"Recherche d'informations pour : {address}")

    try:
        lat, lon = geocode_address(address)
        print(f"Coordonnées : {lat:.6f}, {lon:.6f}")
        data = search_nearby(lat, lon)
        infos = extract_info(data, address)

        if not infos:
            print("Aucune information trouvée sur OpenStreetMap.")
            return

        print("\n=== Résultats ===")
        for i, info in enumerate(infos, 1):
            print(f"\n[{i}] {info['nom']}")
            print(f" Type        : {info['type']}")
            print(f" Téléphone   : {info['téléphone']}")
            print(f" Email       : {info['email']}")
            print(f" Site web    : {info['site']}")
            print(f" Horaires    : {info['horaires']}")
            print(f" Adresse OSM : {info['adresse_OSM']}")
            if info["autres_tags"]:
                print(f" Autres tags : {json.dumps(info['autres_tags'], ensure_ascii=False, indent=2)}")

    except Exception as e:
        print(f"Erreur : {e}")

if __name__ == "__main__":
    main()

