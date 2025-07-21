-- Enable UUID extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Institution table
CREATE TABLE Institution (
    instId UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    instCode VARCHAR(50) NOT NULL,
    instName VARCHAR(100) NOT NULL,
    instAddress TEXT,
    instType VARCHAR(50),
    registeredAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    lastUpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Institution credentials table
CREATE TABLE Institution_Creds (
    id SERIAL PRIMARY KEY
    instId UUID REFERENCES Institution(instId),
    clientKey TEXT,
    clientSecret TEXT,
    registeredAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    lastUpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Member table
CREATE TABLE Member (
    memberId UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nik VARCHAR(20) UNIQUE NOT NULL,
    memberName VARCHAR(100) NOT NULL,
    memberOf UUID REFERENCES Institution(instId),
    beneficiaryIdentifier VARCHAR(50),
    registeredAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    lastUpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Vehicle table
CREATE TABLE Vehicle (
    vehicleId UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    type VARCHAR(50) NOT NULL,
    brand VARCHAR(50) NOT NULL,
    registeredAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    lastUpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Vehicle_Own table (junction table for ownership)
CREATE TABLE Vehicle_Own (
    memberId UUID REFERENCES Member(memberId),
    vehicleId UUID REFERENCES Vehicle(vehicleId),
    ownerType VARCHAR(50) NOT NULL,
    lastUpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (memberId, vehicleId)
);

-- Vehicle_Registrations table
CREATE TABLE Vehicle_Registrations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    vehicleId UUID REFERENCES Vehicle(vehicleId),
    frontViewUrl TEXT,
    rearViewUrl TEXT,
    sideViewUrl TEXT,
    registeredAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    lastUpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);