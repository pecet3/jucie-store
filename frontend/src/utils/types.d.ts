export type Product = {
    id: number
    name: string
    image_url: string
    description: string
}

export type Price = {
    id: number
    price: number
    capacity: number
}

export type BasketItem = {
    productId: number
    productName: string
    capacity: number
    price: number
    strength: number
}