async function createMenu() {
    try {
        const res = await fetch('https://filkom-api.dvnnfrr.my.id/api/v1/menus', {
        method: 'POST',
        headers: {
            'x-api-key': 'Key ',
            'Authorization': 'Bearer '
        },
        body: JSON.stringify({
            "menu_name": "Katsu BBQ",
            "shop_id": "MDE5M2E1MDAtODcyMi1jM2UzLWMwMmEtMzBhOTk5YzIyOGEw",
            "menu_price": 13000,
            "menu_status": "Ada"
        })
    })

    const data = await res.json()

    console.log(data)
    } catch (err) {
        console.error(err)
    }
    
}

createMenu()