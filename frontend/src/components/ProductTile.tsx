import { FaShoppingBasket } from "react-icons/fa"
import { Product } from "../utils/types"
import { Key } from "react"

export const ProductTile = ({ product }: { product: Product }) => {
    return (
        <>
            <h4 className="font-semibold capitalize text-3xl step-title text-white">{product.name}</h4>
            <img src={product.imageUrl} className="w-72 h-72 object-contain" />
            <button className="bg-purple-950 text-white rounded-3xl py-3 px-8 border-4 border-white hover:bg-transparent hover:border-purple-950 hover:text-white duration-300 flex items-center justify-center mt-4">
                <FaShoppingBasket size={32} />
            </button>
        </>
    )
}