-- +goose Up
-- CREATE OR REPLACE FUNCTION delete_old_bridge_request()
-- RETURNS TRIGGER AS $TRIGGER$
-- BEGIN
-- DELETE FROM bridge_requests
-- WHERE created_at < (NOW() - INTERVAL '30 minutes');
-- RETURN NEW;
-- END;
-- $TRIGGER$ LANGUAGE plpgsql;

-- CREATE TRIGGER trigger_delete_old_bridge_request
--     AFTER INSERT ON bridge_requests
--     EXECUTE PROCEDURE delete_old_bridge_request();

-- +goose Down
-- DROP TRIGGER IF EXISTS trigger_delete_old_bridge_request;
-- DROP FUNCTION IF EXISTS delete_old_bridge_request();
