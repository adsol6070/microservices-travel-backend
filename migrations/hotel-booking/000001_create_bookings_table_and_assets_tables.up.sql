CREATE TYPE booking_status AS ENUM (
    'pending',                -- Booking has been created, but not yet confirmed
    'confirmed',              -- Booking is confirmed and reserved
    'awaiting_payment',       -- Booking is confirmed but awaiting payment
    'payment_received',       -- Payment has been received, booking is secure
    'cancelled',              -- Booking has been cancelled by the guest or the hotel
    'cancelled_by_guest',     -- Booking cancelled by the guest
    'cancelled_by_hotel',     -- Booking cancelled by the hotel (e.g., due to overbooking)
    'checked_in',             -- Guest has checked in to the hotel
    'checked_out',            -- Guest has checked out of the hotel
    'no_show',                -- Guest did not show up and failed to cancel
    'awaiting_confirmation',  -- Booking awaiting confirmation from external systems or hotel management
    'pending_review',         -- Booking is pending guest review
    'reviewed',               -- Guest has submitted a review for the booking
    'expired',                -- Booking has expired (e.g., unpaid and the time window has passed)
    'in_progress',            -- Booking is in progress (guest is currently staying)
    'under_payment_review',   -- Payment review process (e.g., fraud check)
    'under_cancellation_review', -- Cancellation request is being reviewed (e.g., by the hotel or the guest)
    'refunded',               -- Payment has been refunded to the guest
    'disputed'                -- Booking is under dispute between guest and hotel
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),   -- Unique booking ID
    asset_id INT NOT NULL,                -- Asset ID for the room or property
    user_id INT NOT NULL,                 -- User ID for the person who made the booking
    status booking_status NOT NULL,       -- Booking status (pending, confirmed, cancelled, etc.)
    start_date DATE NOT NULL,             -- Start date of the booking (check-in date)
    end_date DATE NOT NULL,               -- End date of the booking (check-out date)
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the booking was created
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the booking was last updated

    room_id UUID,                         -- Room ID for the specific room (if applicable)
    room_type VARCHAR(255),               -- Room type (e.g., single, double, suite)
    guest_count INT NOT NULL,             -- Number of guests included in the booking
    special_requests TEXT,                -- Any special requests or notes from the guest

    total_price DECIMAL(10, 2) NOT NULL,  -- Total price of the booking
    payment_status VARCHAR(255),          -- Payment status (e.g., paid, pending, refunded)
    payment_method VARCHAR(255),          -- Payment method used for the booking
    currency VARCHAR(3) NOT NULL,         -- Currency code (e.g., USD, EUR, GBP)
    exchange_rate DECIMAL(10, 4),         -- Exchange rate used for currency conversion
    total_price_local DECIMAL(10, 2),     -- Total price in local currency (if applicable)

    checkin_date DATE,                    -- Actual check-in date
    checkout_date DATE,                   -- Actual check-out date
    is_checkin_flexible BOOLEAN DEFAULT FALSE,  -- If check-in date is flexible
    is_checkout_flexible BOOLEAN DEFAULT FALSE, -- If check-out date is flexible

    guest_name VARCHAR(255),              -- Name of the primary guest
    guest_email VARCHAR(255),             -- Email address of the primary guest
    guest_phone_number VARCHAR(15),       -- Phone number of the primary guest
    guest_nationality VARCHAR(100),       -- Nationality of the primary guest

    hotel_id UUID NOT NULL,                -- Foreign key to the hotel table
    branch_id INT,                        -- Foreign key to the branch table (if applicable)
    booking_channel VARCHAR(100),         -- Channel through which booking was made (website, app, third-party, etc.)

    cancellation_reason TEXT,             -- Reason for cancellation, if applicable
    cancellation_date TIMESTAMP,          -- Date of cancellation, if applicable

    rating INT CHECK (rating >= 1 AND rating <= 5),  -- Rating given by the guest (1 to 5)
    review TEXT,                          -- Guest review text

     -- Additional fields for API integration
    external_booking_id VARCHAR(255),       -- Store the booking ID from the external API
    external_data JSONB,                    -- Store additional data from the external API
    external_sync_attempts INT DEFAULT 0,   -- Number of attempts to sync with external API
    external_sync_last_attempt TIMESTAMP,   -- Timestamp of the last sync attempt with external API
    external_sync_error TEXT,               -- Error message from the last sync attempt

    api_sync_status VARCHAR(50),            -- Status of synchronization (e.g., "pending", "successful", "failed")
    external_payment_status VARCHAR(50),    -- External API's payment status (if needed)
    external_channel VARCHAR(100),          -- Channel provided by the external API (if different)
    api_last_synced TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of last sync with external API

    CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE,  -- Foreign Key for Room
);

-- Create a table for rooms (optional, if not already present)
CREATE TABLE rooms (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),    -- Room ID as UUID
    hotel_id UUID NOT NULL,                            -- Foreign key to the hotels table
    room_type VARCHAR(50),                -- Type of room (Single, Double, Suite, etc.)
    capacity INT NOT NULL,                -- Capacity (max number of guests)
    price DECIMAL(10, 2),                 -- Price per night for the room
    available BOOLEAN DEFAULT TRUE,       -- Whether the room is available
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the room was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the room was last updated

    CONSTRAINT fk_hotel FOREIGN KEY (hotel_id) REFERENCES hotels(id) ON DELETE CASCADE -- Foreign Key for Hotel
);