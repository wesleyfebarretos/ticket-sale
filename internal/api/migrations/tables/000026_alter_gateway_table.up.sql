ALTER TABLE 
    fin.gateway 
ADD COLUMN 
    gateway_provider_id INT NOT NULL UNIQUE 
REFERENCES
    fin.gateway_provider(id);
