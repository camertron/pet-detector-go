package main

type GameResult struct {
    Game_id string `json:"game_id"`
    From_extension bool `json:"from_extension"`
    Game_result struct {
        Score int `json:"score"`
        Game_result_data struct {
            Version string `json:"version"`
            Default_locale_id string `json:"default_locale_id"`
            Trial_csv string `json:"trial_csv"`
            Pauses string `json:"pauses"`
            Locale_id string `json:"locale_id"`
            Tutorial_start int64 `json:"tutorial_start"`
            Num_correct int `json:"num_correct"`
            Num_total int `json:"num_total"`
            Round_csv string `json:"round_csv"`
            Num_trials int `json:"num_trials"`
            Num_defocus int `json:"num_defocus"`
            Split_tests struct {
                Flash_tb_leaderboard string `json:"flash_tb_leaderboard"`
                Flash_hh_difficulty string `json:"flash_hh_difficulty"`
                Flash_mp_score string `json:"flash_mp_score"`
                Flash_wbr_cooperative string `json:"flash_wbr_cooperative"`
                Flash_ec_level_up_time string `json:"flash_ec_level_up_time"`
                Flash_tot_daily_challenge string `json:"flash_tot_daily_challenge"`
                Flash_continuum_clarity string `json:"flash_continuum_clarity"`
                Flash_wbr_stem string `json:"flash_wbr_stem"`
                Flash_rd_equations string `json:"flash_rd_equations"`
                Flash_contextual_length string `json:"flash_contextual_length"`
                Flash_spaced_repetition string `json:"flash_spaced_repetition"`
                Flash_cc_task_feedback string `json:"flash_cc_task_feedback"`
                Flash_tot_new_fit_tutorial string `json:"flash_tot_new_fit_tutorial"`
                } `json:"split_tests"`
            Num_pauses int `json:"num_pauses"`
            Tutorial_finish int64 `json:"tutorial_finish"`
            Flash_game_id string `json:"flash_game_id"`
            Start_level int `json:"start_level"`
            Time int `json:"time"`
            Playtime int `json:"playTime"`
            } `json:"game_result_data"`
        Session_level int `json:"session_level"`
        User_level int `json:"user_level"`
        } `json:"game_result"`
    Mode string `json:"mode"`
    }
