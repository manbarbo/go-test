CREATE TABLE IF NOT EXISTS stock_information (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticker TEXT,
    target_from TEXT,
    target_to TEXT,
    company TEXT,
    action TEXT,
    brokerage TEXT,
    rating_from TEXT,
    rating_to TEXT,
    time TEXT
);