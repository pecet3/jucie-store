import React, { createContext, useContext, useState, ReactNode } from 'react';
import { Product } from './types';


type StoreContextType = {
    products: Product[]
    addProduct: (product: Product) => void
    removeProduct: (name: string) => void
}

const StoreContext = createContext<StoreContextType | undefined>(undefined);

export const StoreProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [products, setProducts] = useState<Product[]>([]);

    const addProduct = (product: Product) => {
        setProducts(prevProducts => [...prevProducts, product]);
    };

    const removeProduct = (name: string) => {
        setProducts(prevProducts => prevProducts.filter(product => product.name !== name));
    };

    return (
        <StoreContext.Provider value={{ products, addProduct, removeProduct }}>
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