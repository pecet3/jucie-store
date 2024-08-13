import { useStoreContext } from "../utils/storeContext";
import { ProductTile } from "../components/ProductTile";

export const Products = () => {
    const { products, prices } = useStoreContext()
    return (
        <div className="flex-grow flex flex-col items-center justify-center py-16">
            <h2 className="text-6xl font-bold text-white mb-12 step-title">ULTIMATE</h2>
            <div className="w-full max-w-screen-xl px-8">
                <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-12">
                    {products.map(prod => (
                        <div key={prod.name} className="border-4 border-white rounded-lg flex flex-col items-center justify-center shadow-lg p-6 min-w-[18rem]">
                            <ProductTile product={prod} />
                        </div>
                    ))}
                </div>
            </div>
        </div >
    )
}