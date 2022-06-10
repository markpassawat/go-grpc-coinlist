-- struct coins
CREATE TABLE IF NOT EXISTS public.coins 
(
    coin_id VARCHAR, 
    symbol VARCHAR, 
    name VARCHAR, 
    image VARCHAR, 
    current_price FLOAT, 
    market_cap_rank INT, 
    create_at TIMESTAMPTZ, 
    update_at TIMESTAMPTZ, 
);
