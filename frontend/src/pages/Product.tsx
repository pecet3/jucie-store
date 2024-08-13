import { useStoreContext } from "../utils/storeContext";
import { ProductTile } from "../components/ProductTile";
import { useParams } from "react-router-dom";

export const Product = () => {
    const { id } = useParams();
    const { getProductById } = useStoreContext()
    const product = getProductById(parseInt(id!))
    console.log(product)
    return (
        <div className="text-2xl">
            {product?.name}
        </div >
    )
}