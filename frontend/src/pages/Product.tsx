import { useStoreContext } from "../utils/storeContext";
import { useParams } from "react-router-dom";
import { FaShoppingBasket } from "react-icons/fa";

export const Product = () => {
    const { id } = useParams();
    const { getProductById } = useStoreContext()
    const product = getProductById(Number(id))
    return (
        <div className="flex-grow flex items-center justify-center py-16 relative">
            <div className="w-full max-w-4xl px-4 relative">
                <div className="absolute -top-14 w-full flex justify-center">
                    <h4 className="font-semibold capitalize text-6xl step-title text-white px-4 py-2">{product?.name}</h4>
                </div>
                <div className="border-4 border-white rounded-lg flex items-center justify-between shadow-lg p-6 mt-8">
                    <img src={product?.image_url} className="w-72 h-72 object-contain" />
                    <div className="flex flex-col justify-center">
                        <p className="text-lg text-white mb-4 text-left max-w-xs">
                            <span className="text-purple-950 text-2xl font-bold">{product?.name}</span>
                            {product?.description}
                        </p>
                        <div className="flex space-x-4 items-center">
                            <div className="">
                                <select id="strength"
                                    className="border-solid bg-black border-white border-4 px-5 py-2 rounded cursor-pointer font-bold w-[200px] bg-transparent">
                                    <option value="" disabled selected hidden>STRENGTH</option>
                                    <option value="3">3 MG</option>
                                    <option value="6">6 MG</option>
                                    <option value="12">12 MG</option>
                                    <option value="18">18 MG</option>
                                </select>
                            </div>

                            <div className="">
                                <select id="size" className="border-solid bg-black border-white border-4 px-5 py-2 rounded cursor-pointer font-bold w-[200px] bg-transparent">
                                    <option value="" disabled selected hidden>SIZE</option>
                                    <option value="30">30 ML</option>
                                    <option value="60">60 ML</option>
                                    <option value="100">100 ML</option>
                                </select>
                            </div>
                            <button className="bg-purple-950 rounded-3xl py-3 px-4 hover:bg-transparent hover:border-purple-950 hover:text-white duration-300 hover:border border border-transparent flex items-center justify-center">
                                <FaShoppingBasket size={24} />
                            </button>
                        </div>
                        <div className="text-center">0</div>

                    </div>
                </div>

            </div>
        </div>
    )
}