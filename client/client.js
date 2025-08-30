$(function() {
	// Create event
	$('#create-event-form').on('submit', function(e) {
		e.preventDefault();
		const name = $(this).find('[name="name"]').val();
		const description = $(this).find('[name="description"]').val();
		const location = $(this).find('[name="location"]').val();
		const datetimeLocal = $(this).find('[name="datetime"]').val();
		// Convert datetime-local value to ISO8601 with seconds and timezone
		let dateTime = "";
		if (datetimeLocal) {
			// datetimeLocal is in format "YYYY-MM-DDTHH:mm"
			const dt = new Date(datetimeLocal);
			dateTime = dt.toISOString(); // "YYYY-MM-DDTHH:mm:ss.sssZ"
		}
		const data = {
			name,
			description,
			location,
			dateTime
		};
		console.log('Creating event:', data);
		$.ajax({
			url: '/events',
			method: 'POST',
			contentType: 'application/json',
			data: JSON.stringify(data),
			success: function(res) {
				console.log('Create event response:', res);
				$('#create-event-result').text(res.message + ' (ID: ' + res.event_id + ')');
			},
			error: function(xhr) {
				console.log('Create event error:', xhr);
				$('#create-event-result').text(xhr.responseJSON ? xhr.responseJSON.message : 'Error creating event');
			}
		});
	});

	// Load events
	$('#load-events').on('click', function() {
		console.log('Loading events...');
		$.get('/events', function(events) {
			console.log('Events loaded:', events);
			const $tbody = $('#events-table tbody');
			$tbody.empty();
			if (Array.isArray(events) && events.length) {
				events.forEach(function(ev) {
					$tbody.append('<tr>' +
						'<td>' + ev.id + '</td>' +
						'<td>' + ev.name + '</td>' +
						'<td>' + ev.description + '</td>' +
						'<td>' + ev.location + '</td>' +
						'<td>' + ev.dateTime + '</td>' +
					'</tr>');
				});
			} else {
				$tbody.append('<tr><td colspan="5">No events found.</td></tr>');
			}
		}).fail(function(xhr) {
			console.log('Load events error:', xhr);
			$('#events-table tbody').html('<tr><td colspan="5">Error loading events</td></tr>');
		});
	});
});
