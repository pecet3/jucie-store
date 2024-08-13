import { FaShoppingBasket } from "react-icons/fa";
import { useStoreContext } from "../utils/storeContext";
import { Product } from "../utils/types";
import { ProductTile } from "../components/ProductTile";
import { useEffect } from "react";

export const Products = () => {
    const { products, addProduct } = useStoreContext()
    const prod: Product = {
        name: "test",
        description: "lorem ipsum",
        imageUrl: "/images/ONI.png"
    }
    useEffect(() => {
        addProduct(prod)
    }, [])
    console.log(products)

    return (
        <div className="flex-grow flex flex-col items-center justify-center py-16">
            <h2 className="text-6xl font-bold text-white mb-12 step-title">ULTIMATE</h2>
            <div className="w-full max-w-screen-xl px-8">
                <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-12">
                    {products.map(prod => {
                        return <ProductTile key={prod.imageUrl} product={prod} />
                    })}
                </div>
            </div>
        </div>
    )
}