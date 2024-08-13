import { FaShoppingBasket } from "react-icons/fa"
import { Product } from "../utils/types"
import { Key } from "react"
import { Link } from "react-router-dom"

export const ProductTile = ({ product }: { product: Product }) => {

    return (
        <>
            <h4 className="font-semibold capitalize text-3xl step-title text-white">{product.name}</h4>
            <img src={`http://127.0.0.1:8090${product.image_url}`} className="w-72 h-72 object-contain" />
            <Link to={`/juices/${product.id}`}
                className="bg-purple-950 text-white rounded-3xl py-3 px-8 border-4
             border-white hover:bg-transparent
              hover:border-purple-950 hover:text-white duration-300 
              flex items-center justify-center mt-4">
                <FaShoppingBasket size={32} />
            </Link>
        </>
    )
}