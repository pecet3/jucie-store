import React, { createContext, useContext, useState, ReactNode } from 'react';
import { BasketItem, Price, Product } from './types';


type StoreContextType = {
    products: Product[]
    addProducts: (products: Product[]) => void
    getProductById: (id: number) => Product | undefined
    prices: Price[]
    addPrices: (prices: Price[]) => void
}

const StoreContext = createContext<StoreContextType | undefined>(undefined);

export const StoreProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [products, setProducts] = useState<Product[]>([{
        id: 0,
        name: "",
        image_url: "",
        description: "",
    }]);
    const addProducts = (products: Product[]) => {
        setProducts(products);
    };
    const getProductById = (id: number) => {
        return products.find(p => p.id === id)
    }

    const [prices, setPrices] = useState<Price[]>([{
        id: 0,
        price: 0,
        capacity: 0
    }])
    const addPrices = (prices: Price[]) => {
        setPrices(prices)
    }

    const [basket, setBasket] = useState<BasketItem[]>([])


    return (
        <StoreContext.Provider value={{ products, addProducts, getProductById, prices, addPrices }}>
            {children}
        </StoreContext.Provider>
    );
};

export const useStoreContext = () => {
    const context = useContext(StoreContext);
    if (context === undefined) {
        throw new Error('useStoreContext must be used within a StoreProvider');
    }
    return context;
};