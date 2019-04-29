from op5beat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Op5beat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        op5beat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("op5beat is running"))
        exit_code = op5beat_proc.kill_and_wait()
        assert exit_code == 0
