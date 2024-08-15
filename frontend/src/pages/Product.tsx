import { useStoreContext } from "../utils/storeContext";
import { useParams } from "react-router-dom";
import { FaShoppingBasket } from "react-icons/fa";
import { useEffect, useState } from "react";
import { BasketItem, Price } from "../utils/types";
import toast from "react-hot-toast";

export const Product = () => {
    const { id } = useParams();
    const { getProductById, prices } = useStoreContext();
    const product = getProductById(Number(id));
    const [strength, setStrength] = useState("3");
    const [price, setPrice] = useState(prices[0].price.toString())
    const handleButtonClick = () => {
        if (!strength) return;
        const capacity = prices.find(p => p.price === Number(price))?.capacity
        if (!capacity) return;

        const basketItem: BasketItem = {
            productId: Number(id),
            productName: product!.name,
            strength: Number(strength),
            capacity: capacity,
            price: Number(price)
        }
        console.log(basketItem)

        toast.success(`Added a ${product?.name}`)
    };

    return (
        <div className="flex-grow flex items-center justify-center py-16 relative">
            <div className="w-full max-w-4xl px-4 relative">
                <div className="absolute -top-14 w-full flex justify-center">
                    <h4 className="font-semibold capitalize text-6xl step-title text-white px-4 py-2">{product?.name}</h4>
                </div>
                <div className="border-4 border-white rounded-lg flex items-center justify-between shadow-lg p-6 mt-8">
                    <img src={product?.image_url} className="w-72 h-72 object-contain" alt={product?.name} />
                    <div className="flex flex-col justify-center">
                        <p className="text-lg text-white mb-4 text-left max-w-xs">
                            <span className="text-purple-950 text-2xl font-bold">{product?.name}</span>
                            {product?.description}
                        </p>
                        <div className="flex space-x-4 items-center">
                            <div className="">
                                <select
                                    id="strength"
                                    className="border-solid bg-black border-white border-4 px-5 py-2 rounded cursor-pointer font-bold w-[200px] bg-transparent"
                                    value={strength}
                                    onChange={(e) => setStrength(e.target.value)}
                                >
                                    <option value="3" className="bg-black">3 MG</option>
                                    <option value="6" className="bg-black">6 MG</option>
                                    <option value="12" className="bg-black">12 MG</option>
                                    <option value="18" className="bg-black">18 MG</option>
                                </select>
                            </div>
                            <div className="">
                                <select
                                    id="capacity"
                                    className="border-solid bg-black border-white border-4 px-5 py-2 rounded cursor-pointer font-bold w-[200px] bg-transparent"
                                    value={price}
                                    onChange={(e) => setPrice(e.target?.value)}
                                >
                                    {prices?.map(p => (
                                        <option key={p.id} value={p.price} className="bg-black">{p.capacity} ML</option>
                                    ))}
                                </select>
                            </div>
                            <button
                                className="bg-purple-950 rounded-3xl py-3 px-4 hover:bg-transparent hover:border-purple-950 hover:text-white duration-300 hover:border border border-transparent flex items-center justify-center"
                                onClick={handleButtonClick}
                            >
                                <FaShoppingBasket size={24} />
                            </button>
                        </div>
                        <div className="text-center text-2xl font-mono mr-12 mt-4">{price} PLN</div>
                    </div>
                </div>
            </div>
        </div>
    );
};